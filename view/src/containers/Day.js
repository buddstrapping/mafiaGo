import React from 'react';
import logo from '../assets/logo.svg';
import '../assets/App.css';
import swal from "sweetalert";
import '../assets/LinearLayoutComponent.css'
import Switch from '@material-ui/core/Switch';
import PaymentCard from 'react-payment-card-component'
import axios from 'axios'
import Button from '@material-ui/core/Button';

class Day extends React.Component {
    constructor(props) {
        super(props);
        this.state = ({
            name: this.props.location.state.Name,
            career: "배정중..",
            filpButton: false,
            flipped: false
        });
    }

    handleChange = name => event => {
        this.setState({ ...this.state, [name]: event.target.checked });
        if (!this.state.filpButton && this.state.career === "배정중..") {
            axios.post('/day/check', {
                name: this.state.name,
                career: this.state.career,
                state : 'O'
            }).then(res => {
                this.setState({career: res.data});
                this.flipCard();
            })

            if(!this.state.flipped)
                this.flipCard();
        }
        else
            this.flipCard()
    }

    flipCard() {
        const flipped = !this.state.flipped
        this.setState({ flipped })
    }
    
    gogo = () => {
        swal({
            text: "Test all of message",
            icon: "warning",
            buttons: false
        })

        axios.post('/day/countdown', {
            name: this.state.name,
            career: this.state.career,
            target : 'O'
        }).then(res => {
            swal({
                text: "Good!",
                icon: "success"
            })
        })
    }

    render() {
        return (

            <div className="App">
                <header className="App-header">
                    <img src={logo} className="App-logo" alt="logo" />
                    <p>
                        Edit <code>src/App.js</code> and save to reload.
              </p>
                    <Switch
                        onChange={this.handleChange('filpButton')}
                        value="night"
                        color="default"
                        inputProps={{ 'aria-label': 'Night Switch' }}
                    />

                    <PaymentCard
                        bank="santander"
                        model="normal"
                        type="black"
                        number=""
                        brand="mastercard"
                        cvv={this.state.career}
                        holderName={this.state.name}
                        expiration="12/20"
                        flipped={this.state.flipped}
                    />
                    <p></p>
                    <Button variant="contained" color="primary" onClick={this.gogo}>GO</Button>
                </header>
            </div>

        );
    };
}

export default Day;