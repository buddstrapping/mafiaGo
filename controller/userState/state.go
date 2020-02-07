package userState

import (
	"fmt"
	"io/ioutil"
	"mafiaGo/model"
	"math/rand"
	"time"

	"encoding/json"
)

var set model.VarSet
var offline, online []model.User
var resLen int
var liveAll, polBool bool
var resMessage string = "취합중..."
var popularPerson map[string]int = make(map[string]int)
var maxIdx []string
var gaming bool = false

func LoadOfflineList() {
	b, err := ioutil.ReadFile("./model/user_list.json")

	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	json.Unmarshal(b, &offline)

	set.AllNum = 11
	set.MafiaNum = 3
	set.DocNum = 1
	set.PolNum = 2

	fmt.Println(offline)
}

/* 게임 초기화: 온라인 유저 수로 게임을 진행하기 때문에 이를 초기화 */
func ResetGame() string {
	online = nil
	maxIdx = nil
	dup = nil
	popularPerson = make(map[string]int)
	resMessage = "취합중..."
	gaming = false
	resLen = 0
	liveAll = false
	polBool = false

	set.AllNum = 11
	set.MafiaNum = 3
	set.DocNum = 1
	set.PolNum = 2

	fmt.Printf("[RESET]: Complete\n")
	return "success"
}

/* 직업 수 설정 */
func SetGameRule(s model.VarSet) string {
	if len(online) == 0 {
		set = s
		debugCareer()
		return "success"
	}
	return "fail"
}

/* 접속한 유저가 유효한지, 이미 준비상태인지 체크 */
func UpdateOnline(newUser model.User) int {
	// check vaild user
	check := false
	for _, user := range offline {
		if newUser.Name == user.Name {
			check = true
			break
		}
	}

	if check {
		if CheckResume(newUser) {
			return 201
		} else if gaming {
			return 401
		}

		online = append(online, newUser)
		debugCareer()
		CheckToAllReady()
		return 200
	}

	return 400
}

func CheckResume(newUser model.User) bool {
	for i, user := range online {
		if newUser.Name == user.Name {
			if online[i].State == "X" {
				online[i].State = "O"
				set.AllNum++
			}
			return true
		}
	}
	return false
}

func CheckToAllReady() {
	if len(online) == set.AllNum {
		PickCareer()
		gaming = true
	}
}

/* 직업 랜덤 배치 */
func PickCareer() {
	var picked []int
	var i = 0
	for i < set.MafiaNum {
		timeSource := rand.NewSource(time.Now().UnixNano())
		random := rand.New(timeSource)
		pick := random.Intn(set.AllNum - 1)
		fmt.Printf("MAFIA PICK:%d\n", pick)
		check := false
		for _, j := range picked {
			if j == pick {
				check = true
			}
		}

		if !check {
			online[pick].Career = "마피아"
			picked = append(picked, pick)
			i++
		}
	}

	i = 0
	for i < set.DocNum {
		timeSource := rand.NewSource(time.Now().UnixNano())
		random := rand.New(timeSource)
		pick := random.Intn(set.AllNum - 1)
		fmt.Printf("DOC PICK:%d\n", pick)
		check := false
		for _, j := range picked {
			if j == pick {
				check = true
			}
		}

		if !check {
			online[pick].Career = "의사"
			picked = append(picked, pick)
			i++
		}
	}

	i = 0
	for i < set.PolNum {
		timeSource := rand.NewSource(time.Now().UnixNano())
		random := rand.New(timeSource)
		pick := random.Intn(set.AllNum - 1)
		fmt.Printf("POL PICK:%d\n", pick)
		check := false
		for _, j := range picked {
			if j == pick {
				check = true
			}
		}

		if !check {
			online[pick].Career = "경찰"
			picked = append(picked, pick)
			i++
		}
	}

	debugCareer()

}

func debugCareer() {
	for _, i := range online {
		fmt.Printf("[RES] %s is %s in %s\n", i.Name, i.Career, i.Target)
	}
	fmt.Printf("[SET] all : %d mafia: %d doc: %d pol: %d\n", set.AllNum, set.MafiaNum, set.DocNum, set.PolNum)
}

func CheckCareer(t model.User) string {
	if len(online) < set.AllNum {
		return "배정중.."
	}

	for _, i := range online {
		if i.Name == t.Name {
			return i.Career
		}
	}

	return "Error"
}

func LoadPeople() []model.User {
	var liveList []model.User
	for _, i := range online {
		if i.State == "O" {
			liveList = append(liveList, i)
		}
	}

	return liveList
}

var dup []string

func ProcessResult(t model.User) int {
	/* 죽은 애 클릭했는지 여부 */
	for _, i := range online {
		if i.Name == t.Target {
			if i.State == "X" {
				return 202
			}
			break
		}
	}
	/* 중복 요청 체크 */
	for _, i := range dup {
		if i == t.Name {
			return 201
		}
	}
	dup = append(dup, t.Name)

	resLen++
	polBool = false
	fmt.Printf("USER: %s => %s\n", t.Name, t.Target)

	for i := range online {
		if online[i].Name == t.Name {
			online[i].Target = t.Target
			break
		}
	}

	if resLen == set.AllNum {
		makeResult()
		resLen = 0
		dup = nil
		popularPerson = make(map[string]int)
		maxIdx = nil
	}

	return 200
}

func makeResult() {
	var mTarget string
	var dTarget []string
	for _, i := range online {
		if i.State == "O" {
			fmt.Printf("[RESULT] %s => %s\n", i.Name, i.Target)
			switch {
			case i.Career == "마피아":
				mTarget = i.Target
			case i.Career == "의사":
				dTarget = append(dTarget, i.Target)
			case i.Career == "경찰":
				i.Target = search(i.Target)
			default:
				popularPerson[i.Target]++
				if len(maxIdx) == 0 {
					maxIdx = append(maxIdx, i.Target)
				} else if popularPerson[i.Target] > popularPerson[maxIdx[0]] {
					maxIdx = append([]string(nil), i.Target)
				} else if popularPerson[i.Target] == popularPerson[maxIdx[0]] {
					if i.Target == maxIdx[0] {
						maxIdx = append([]string(nil), i.Target)
					} else {
						maxIdx = append(maxIdx, i.Target)
					}

				}
			}
		}
	}

	isLive := false
	for _, i := range dTarget {
		if i == mTarget {
			isLive = true
			break
		}
	}

	liveAll = false
	resMessage = "[알림]: "
	if isLive {
		resMessage += "전원 생존!\n"
		liveAll = true
	} else {
		resMessage += mTarget + "이(가) 사-망\n"
	}

	resMessage += "[우수시민]: "
	for _, str := range maxIdx {
		resMessage += str + ", "
	}

	maxIdx = nil

	resMessage = resMessage[:len(resMessage)-2]

	polBool = true
}

func CheckRes(user model.User) (string, bool) {
	searchRes := ""
	fmt.Printf("C:%s\n", user.Career)
	if polBool {
		if user.Career == "경찰" {
			for _, i := range online {
				if user.Name == i.Name {
					searchRes = search(i.Target)
				}
			}
		}
	}
	return resMessage + searchRes + "\n", liveAll
}

func search(name string) string {
	for _, i := range online {
		if i.Name == name && i.Career == "마피아" {
			return " YES"
		}
	}
	return " NO"
}

func DeadRequest(user model.User) string {
	for i, _ := range online {
		if online[i].Name == user.Name {
			if online[i].State == "O" {
				online[i].State = "X"
				set.AllNum--
				gaming = Victory()
			}
			return online[i].State
		}
	}

	return "Error"
}

func Victory() bool {
	/* 마피아가 다 죽어있는지 확인 */
	res := false
	for i, _ := range online {
		if online[i].Career == "마피아" && online[i].State == "O" {
			res = true
			break
		}
	}

	if !res {
		return false
	}

	var nCnt, mCnt int
	/* 과반수 확인 */
	for i, _ := range online {
		if online[i].State == "O" {
			if online[i].Career == "마피아" {
				mCnt++
			} else {
				nCnt++
			}
		}
	}

	if mCnt >= nCnt {
		return false
	}

	return true
}
