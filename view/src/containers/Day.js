import React from 'react';
import logo from '../assets/logo.svg';
import '../assets/App.css';
import swal from "sweetalert";
import Switch from '@material-ui/core/Switch';
import PaymentCard from 'react-payment-card-component'
import axios from 'axios'
import Button from '@material-ui/core/Button';
import '../assets/start.scss'
import nightLogo from '../assets/moon.svg'
import { makeStyles, withStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Divider from '@material-ui/core/Divider';
import InboxIcon from '@material-ui/icons/Inbox';
import DraftsIcon from '@material-ui/icons/Drafts';
import audioFile from '../assets/audio.mp3'
import audioDeath from '../assets/death.mp3'

class Day extends React.Component {
    constructor(props) {
        super(props);
        this.state = ({
            name: this.props.location.state.Name,
            career: "배정중..",
            filpButton: false,
            flipped: false,
            date: "",
            rightOpen: false,
            people: []
        });
    }

    componentDidMount() {
        this.getDate();
    }

    toggleSidebar = async (event) => {
        let key = `${event.currentTarget.parentNode.id}Open`;
        this.setState({ [key]: !this.state[key] })

        /* 생존자 리스트 불러오기 */
        if (!this.state.rightOpen) {
            let res = await axios.post('/day/load');
            await this.bindData(res.data);
        }
    }

    myToggle = () => {
        this.setState({
            rightOpen: !this.state.rightOpen
        })
    }


    async bindData(data) {
        let arr = []
        for (let obj of data) {
            arr.push(obj);
        }

        this.setState({
            people: arr
        })
    }

    handleChange = name => event => {
        this.setState({ ...this.state, [name]: event.target.checked });
        if (!this.state.filpButton && this.state.career === "배정중..") {
            axios.post('/day/check', {
                name: this.state.name,
                career: this.state.career,
                state: 'O'
            }).then(res => {
                this.setState({ career: res.data });
                this.flipCard();
            }).catch((error) => {
                this.flipCard();
            })
        }
        else
            this.flipCard()
    }

    flipCard() {
        const flipped = !this.state.flipped
        this.setState({ flipped })
    }

    getResult = () => {
        swal({
            text: "결과 요청중...",
            icon: "warning",
            buttons: false,
            closeOnClickOutside: false
        })

        setTimeout(() => {
            axios.post('/night/checkRes', {
                name: this.state.name,
                career: this.state.career,
            })
                .then(res => {
                    //alert(JSON.stringify(res))
                    if (res.data.liveAll) {
                        this.playAudio();
                    }
                    swal({
                        text: res.data.msg
                    }).then(() => {
                        this.pauseAudio();
                    })
                }).catch((err) => {
                    swal({
                        text: "재요청바람",
                        icon: "error"
                    })
                })
        }, 800)
    }

    playAudio = () => {
        let audio = document.querySelector("audio");
        audio.currentTime = 0;
        audio.src = audio.children[0].src;
        audio.play();

        // setTimeout(() => {
        //     audio.pause();
        // }, 5000);
    }

    pauseAudio = () => {
        let audio = document.querySelector("audio");
        audio.pause();
    }

    requestDead = () => {
        swal({
            text: "실수아니지?",
            buttons: true,
            dangerMode: true
        }).then((willDelete) => {
            if (willDelete) {
                swal({
                    text: "처리중..",
                    icon: "warning",
                    buttons: false,
                    closeOnClickOutside: false
                })

                setTimeout(() => {
                    axios.post('/night/deadRequest', {
                        name: this.state.name,
                        state: "O"
                    })
                        .then(res => {

                            let audio = document.querySelector("audio");
                            audio.currentTime = 0;
                            audio.src = audio.children[1].src;
                            audio.play(); 

                            var str = "";
                            if (res.data == "X") {
                                str = "잘가~"
                            } else {
                                str = "Error"
                            }
                            swal({
                                text: str,
                                closeOnClickOutside: false,
                                buttons: false
                            })

                            setTimeout(() => {
                                this.props.history.push({
                                    pathname: '/'
                                });
                                swal.close();
                            }, 2200)                            
                     
                        }).catch((err) => {
                            swal({
                                text: "재요청바람",
                                icon: "error"
                            })
                        })
                }, 800)
            } else {
                swal({
                    text: "콱씨마!",
                })
            }
        })
    }


    getDate = () => {
        var today = new Date();
        var mm = today.getMonth() + 1;
        var yyyy = today.getFullYear();

        if (mm < 10) {
            mm = '0' + mm;
        }

        today = mm + '/' + yyyy

        this.setState({ date: today });
    }

    render() {
        const BootstrapButton = withStyles({
            root: {
                margin: '10px',
                boxShadow: 'none',
                textTransform: 'none',
                fontSize: 16,
                padding: '6px 12px',
                border: '1px solid', 
                lineHeight: 1.5,
                backgroundColor: '#000000',
                borderColor: '#000000',
                fontFamily: [
                    '-apple-system',
                    'BlinkMacSystemFont',
                    '"Segoe UI"',
                    'Roboto',
                    '"Helvetica Neue"',
                    'Arial',
                    'sans-serif',
                    '"Apple Color Emoji"',
                    '"Segoe UI Emoji"',
                    '"Segoe UI Symbol"',
                ].join(','),
                '&:hover': {
                    backgroundColor: '#0069d9',
                    borderColor: '#0062cc',
                    boxShadow: 'none',
                },
                '&:active': {
                    boxShadow: 'none',
                    backgroundColor: '#0062cc',
                    borderColor: '#005cbf',
                },
                '&:focus': {
                    boxShadow: '0 0 0 0.2rem rgba(0,123,255,.5)',
                },
            },
        })(Button);

        let rightOpen = this.state.rightOpen ? 'open' : 'closed';

        const useStyles = makeStyles(theme => ({
            root: {
                width: '100%',
                maxWidth: 360,
                backgroundColor: theme.palette.background.paper,
            },
        }));

        const buttonStyles = makeStyles(theme => ({
            root: {
                '& > *': {
                    margin: theme.spacing(1),
                },
            },
        }));

        return (
            <div id='layout'>
                <audio id="myAudio">
                    <source src={audioFile} type="audio/mpeg" />
                    <source src={audioDeath} type="audio/mpeg" />
                </audio>
                <div id='main'>
                    <div className="App">
                        <header className="App-header">
                            <img src="https://upload.wikimedia.org/wikipedia/commons/4/45/Logo_Mafia.svg" className="App-logo-Day" alt="logo" />
                            <p>
                            </p>
                            <PaymentCard
                                bank="santander"
                                model="normal"
                                type="black"
                                number="WELCOME_TO_MAFIA"
                                brand="mastercard"
                                cvv={this.state.career}
                                holderName={this.state.name}
                                expiration={this.state.date}
                                flipped={this.state.flipped}
                            />
                            <Switch
                                onChange={this.handleChange('filpButton')}
                                value="night"
                                color="default"
                                inputProps={{ 'aria-label': 'Night Switch' }}
                            />
                            <p></p>
                            <div className={buttonStyles.root}>
                                <BootstrapButton variant="contained" color="primary" onClick={this.getResult}>RESULT</BootstrapButton>
                                <BootstrapButton variant="contained" color="secondary" onClick={this.requestDead}>DEAD</BootstrapButton>
                            </div>

                        </header>
                    </div>
                </div>
                <div id='right' className={rightOpen} >
                    <div className='icon'
                        onClick={this.toggleSidebar} >
                        <img src={nightLogo} className="Option-logo" alt="option" />
                    </div>
                    <div className={`sidebar ${rightOpen}`} >
                        <div className='header'>
                            <h3 className='title'>
                                Night
                                </h3>
                        </div>
                        <div className='content'>
                            <div className={useStyles.root}>
                                {/* <List component="nav" aria-label="main mailbox folders">
                                    {this.state.people.map((con, i) => {
                                        return (<ListItem
                                            button
                                            onClick={this.handleListItemClick(i)}
                                        >
                                            <ListItemIcon>
                                                <DraftsIcon />
                                            </ListItemIcon>
                                            <ListItemText primary={con.name} />
                                        </ListItem>);
                                    })}
                                </List> */}
                                <List component="nav" aria-label="main mailbox folders">
                                    {this.state.people.map((con, i) => {
                                        return (
                                            <ListItem
                                                button
                                                onClick={() => {
                                                    if (this.state.career !== "배정중..") {
                                                        if (con.name !== this.state.name
                                                            || (this.state.career === "시민" || this.state.career === "의사")) {
                                                            swal({
                                                                text: con.name + "이(가) 확실해?",
                                                                buttons: true,
                                                                dangerMode: true
                                                            }).then((willDelete) => {
                                                                if (willDelete) {
                                                                    swal({
                                                                        text: "처리중...",
                                                                        icon: "warning",
                                                                        buttons: false,
                                                                        closeOnClickOutside: false
                                                                    })

                                                                    axios.post("/night/setTarget", {
                                                                        name: this.state.name,
                                                                        career: this.state.career,
                                                                        state: 'O',
                                                                        Target: `${con.name}`
                                                                    }).then((res) => {

                                                                        if (res.status === 200) {
                                                                            setTimeout(() => {
                                                                                swal({
                                                                                    text: "처리 완료",
                                                                                    icon: "success",
                                                                                    buttons: false,
                                                                                    closeOnClickOutside: false
                                                                                });

                                                                            }, 700);
                                                                        } else if (res.status === 201) {
                                                                            setTimeout(() => {
                                                                                swal({
                                                                                    text: "중복 요청 ㅡㅡ*",
                                                                                    icon: "error",
                                                                                    buttons: false,
                                                                                    closeOnClickOutside: false
                                                                                });
                                                                            }, 700);
                                                                        } else if (res.status === 202) {
                                                                            setTimeout(() => {
                                                                                swal({
                                                                                    text: "목록 초기화 해라 ㅡㅡ*",
                                                                                    icon: "error",
                                                                                    buttons: false,
                                                                                    closeOnClickOutside: false
                                                                                });
                                                                            }, 700);
                                                                        }
                                                                    }).catch((error) => {
                                                                        setTimeout(() => {
                                                                            swal({
                                                                                text: "재요청바람",
                                                                                icon: "error",
                                                                                buttons: false,
                                                                                closeOnClickOutside: false
                                                                            });
                                                                        }, 700);
                                                                    });

                                                                    setTimeout(() => {
                                                                        swal.close();
                                                                        this.myToggle();
                                                                    }, 1800);
                                                                }
                                                                else {
                                                                    swal({
                                                                        text: "콱씨마!",
                                                                    })
                                                                }
                                                            })
                                                        }
                                                    }
                                                }}
                                            >
                                                <ListItemIcon>
                                                    <DraftsIcon />
                                                </ListItemIcon>
                                                <ListItemText primary={con.name}></ListItemText>
                                            </ListItem>);
                                    })}
                                </List>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    };
}





class MyList extends React.Component {
    constructor(props) {
        super(props)
        this.state = ({
            selectedIndex: "",
        });
    }


    handleListItemClick(e, i) {
        this.setState({ selectedIndex: i })
        // alert(i +" picked!");
    }

    render() {
        const useStyles = makeStyles(theme => ({
            root: {
                width: '100%',
                maxWidth: 360,
                backgroundColor: theme.palette.background.paper,
            },
        }));

        return (
            <div className={useStyles.root}>
                <List component="nav" aria-label="main mailbox folders">
                    <ListItem
                        button
                        onClick={event => this.handleListItemClick(event, 0)}
                    >
                        <ListItemIcon>
                            <InboxIcon />
                        </ListItemIcon>
                        <ListItemText primary="Inbox" />
                    </ListItem>
                    <ListItem
                        button
                        onClick={event => this.handleListItemClick(event, 1)}
                    >
                        <ListItemIcon>
                            <DraftsIcon />
                        </ListItemIcon>
                        <ListItemText primary="Drafts" />
                    </ListItem>
                </List>
                <Divider />
                <List component="nav" aria-label="secondary mailbox folder">
                    <ListItem
                        button
                        onClick={event => this.handleListItemClick(event, 2)}
                    >
                        <ListItemText primary="Trash" />
                    </ListItem>
                    <ListItem
                        button
                        onClick={event => this.handleListItemClick(event, 3)}
                    >
                        <ListItemText primary="Spam" />
                    </ListItem>
                </List>
            </div>
        );
    }


}





export default Day;