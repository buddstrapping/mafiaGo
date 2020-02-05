package userState

import (
	"fmt"
	"io/ioutil"
	"mafiaGo/model"
	"math/rand"
	"time"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

var set model.VarSet
var offline, online, deadUser []model.User
var queue []*gin.Context

func LoadOfflineList() {
	b, err := ioutil.ReadFile("./model/user_list.json")

	if err != nil {
		fmt.Println("Err:", err)
		return
	}

	json.Unmarshal(b, &offline)

	set.AllNum = 11
	set.MafiaNum = 3
	set.DocNum = 2
	set.PolNum = 1

	fmt.Println(offline)
}

/* 게임 초기화: 온라인 유저 수로 게임을 진행하기 때문에 이를 초기화 */
func ResetGame() {
	online = nil
	deadUser = nil
}

/* 직업 수 설정 */
func SetGameRule(s model.VarSet) {
	set = s
	debugCareer()
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
		} else {
			online = append(online, newUser)
			debugCareer()
			CheckToAllReady()
		}
		return 200
	}

	return 400
}

func CheckResume(newUser model.User) bool {
	for _, user := range online {
		if newUser.Name == user.Name {
			return true
		}
	}
	return false
}

func CheckToAllReady() {
	if len(online) == set.AllNum {
		PickCareer()
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

func insertQueue(c *gin.Context) {
	queue = append(queue, c)
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

func UpdateLive() {

}

func TestAll(c *gin.Context) {
}
