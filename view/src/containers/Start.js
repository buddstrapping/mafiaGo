import React from 'react';
import optionLogo from '../assets/repairing-service.svg'
import '../assets/App.css';
import swal from "sweetalert";
import Button from '@material-ui/core/Button';
import { makeStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import axios from 'axios'
import '../assets/start.scss'


class Start extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            allNum: 11,
            mafiaNum: 3,
            docNum: 2,
            polNum: 1,
            leftOpen: false,
            rightOpen: false
        }
        //setTimeout(() => { props.history.push("/day"); }, 1000);
    }

    SetGame = () => {
        if (parseInt(this.state.allNum) < parseInt(this.state.mafiaNum) + parseInt(this.state.docNum) + parseInt(this.state.polNum)
            || parseInt(this.state.mafiaNum) === 0) {
            swal("인원수 및 마피수 체크")
        }
        else {
            swal({
                text: "요청 중...",
                icon: "warning",
                buttons: false
            })
            axios.post('/start/setNum', {
                allNum: parseInt(this.state.allNum),
                mafiaNum: parseInt(this.state.mafiaNum),
                docNum: parseInt(this.state.docNum),
                polNum: parseInt(this.state.polNum),
            }).then(res => {
                if (res.status === 200) {
                    setTimeout(() => {
                        swal({
                            text: "셋팅완료",
                            icon: "success",
                            buttons: false
                        });
                    }, 700);

                    setTimeout(() => {
                        swal.close();
                    }, 1800);
                }
            }).catch(error => {
                swal({
                    title: "셋팅 실패",
                    icon: "error",
                    button: {
                        cancel: "Close"
                    }
                });
            });
        }
    }

    handleChange = (e) => {
        let value = e.target.value;
        if (!Number(value) && value !== '0') {
            swal("Please insert Number")
        }
        else {
            this.setState({
                [e.target.name]: e.target.value
            });
        }
    }

    toggleSidebar = (event) => {
        let key = `${event.currentTarget.parentNode.id}Open`;
        this.setState({ [key]: !this.state[key] })
    }


    readyToStart = async () => {
        swal({
            text: '이름 입력',
            content: "input",
        }).then((name) => {

            swal({
                text: "요청 중...",
                icon: "warning",
                buttons: false
            })

            /* 디버깅용 redirect */
            this.props.history.push({
                pathname: '/day',
                state: { Name: `${name}` }
            });


            axios.post('/start/ready', {
                Name: `${name}`,
                Career: "시민",
                State: 'O'
            })
                .then(res => {
                    var str = "";
                    if (res.status === 200) {
                        str = "참가완료"
                    } else if (res.status === 201) {
                        str = "재진입"
                    }

                    setTimeout(() => {
                        swal({
                            title: `${str}`,
                            icon: "success",
                        }).then(() => {
                            this.props.history.push({
                                pathname: '/day',
                                state: { Name: `${name}` }
                            });
                        });
                    }, 1000);
                })
                .catch(function (error) {
                    swal({
                        text: "이름 똑바로 입력해라",
                        icon: "error",
                        button: {
                            cancel: "Close"
                        }
                    });
                });
        });
    }

    ResetGame = () => {
        swal({
            text: "리셋 중...",
            icon: "warning",
            buttons: false
        })

        axios.post('/start/reset')
            .then(res => {
                if (res.status === 200) {
                    setTimeout(() => {
                        swal({
                            text: "리셋 완료",
                            icon: "success",
                        });
                    }, 1000);
                }
            })
            .catch(function (error) {
                swal({
                    text: "요청 실패",
                    icon: "error",
                    button: {
                        cancel: "Close"
                    }
                });
            });
    }


    render() {
        let leftOpen = this.state.leftOpen ? 'open' : 'closed';

        const classes = makeStyles(theme => ({
            root: {
                '& > *': {
                    margin: theme.spacing(1),
                    width: 200,
                }
            }
        }));

        return (
            <div id='layout'>
                <div id='left' className={leftOpen} >
                    <div className='icon'
                        onClick={this.toggleSidebar} >
                        <img src={optionLogo} alt="option"/>
                    </div>
                    <div className={`sidebar ${leftOpen}`} >
                        <div className='header'>
                            <h3 className='title'>
                                Settings
                    </h3>
                        </div>
                        <div className='content'>
                            <form className={classes.root} noValidate autoComplete="off">

                                <TextField id="filled-allNum" label="User Number" variant="outlined"
                                    name='allNum' margin='normal' fullWidth={true} onChange={this.handleChange} />

                                <TextField id="filled-mafiaNum" label="Mafia Number" variant="outlined"
                                    name='mafiaNum' margin='normal' fullWidth={true} onChange={this.handleChange} />

                                <TextField id="filled-docNum" label="Doctor Number" variant="outlined"
                                    name='docNum' margin='normal' fullWidth={true} onChange={this.handleChange} />

                                <TextField id="filled-polNum" label="Police Number" variant="outlined"
                                    name='polNum' margin='normal' fullWidth={true} onChange={this.handleChange} />
                                <p>
                                    <Button variant="contained" color="primary"
                                        fullWidth={true} onClick={this.SetGame}>Set</Button>
                                </p>
                                <Button variant="contained" color="secondary"
                                    fullWidth={true} onClick={this.ResetGame}>Reset</Button>

                            </form>
                        </div>
                    </div>
                </div>
                <div id='main'>
                    <div className="App">
                        <header className="App-header">
                            <img src="https://upload.wikimedia.org/wikipedia/commons/4/45/Logo_Mafia.svg" className="App-logo" alt="logo" />
                            <div>
                                <Button variant="contained" color="primary" onClick={this.readyToStart}>Ready</Button>
                            </div>
                        </header>
                    </div>
                </div>
            </div>
        );
    }
}

export default Start;