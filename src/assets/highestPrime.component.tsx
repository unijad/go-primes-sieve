import * as React from "react";
import './highestPrime.style.scss';

export default class GetHighestPrimeFormWidget extends React.Component {
    state = {
        highestPrime: null,
        prime_limit_input: null,
        status: {
            type: null,
            message: null
        }
    }

    getThePrime(ev) {
        ev.preventDefault();

        let {prime_limit_input} = this.state;
        prime_limit_input = Number(prime_limit_input);

        if (isNaN(prime_limit_input) || prime_limit_input <= 2) {
            this.setState({
                status: {
                    type: 'error',
                    message: 'Input must be an integer > 2'
                }
            })
            return;
        }

        fetch('http://localhost:8080/get-previous-prime', {
            method: 'POST',
            body: JSON.stringify({
                "limit": Number(prime_limit_input)
            })
        })
            .then(res => res.json())
            .then(json => {
                const {highestPrime} = json;
                this.setState({
                    highestPrime,
                    status: {
                        type: null,
                        message: null
                    }
                });
            })
        return false;
    }

    primeLimitOnChange(ev) {
        this.setState({
            prime_limit_input: ev.target.value
        })
    }

    clearResult() {
        this.setState({
            highestPrime: null,
            prime_limit_input: null
        })
    }

    renderForm() {
        const {status} = this.state;

        return <form id="prime-form" onSubmit={(ev) => this.getThePrime(ev)}>
            <fieldset id="title">Highest Prime</fieldset>
            {status.type != null
                ? <p id="validation" className={status.type}>{status.message}</p>
                : null
            }
            <input
                className="form-field"
                type="number"
                name="prime_limit_input"
                placeholder={`enter prime limit`}
                onChange={(ev) => this.primeLimitOnChange(ev)}
            />
            <button className="button primary" type={`submit`}>Get Prime</button>
        </form>
    }

    renderResult(highestPrime) {
        return <div id="prime-result">
            <p>Highest Prime is:</p>
            <h2>{highestPrime}</h2>
            <button className="button secondary" onClick={() => this.clearResult()}>Try another Number</button>
        </div>
    }

    render() {
        const {highestPrime} = this.state;

        if (highestPrime) {
            return this.renderResult(highestPrime);
        }

        return this.renderForm()
    }
}