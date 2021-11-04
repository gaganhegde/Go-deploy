import React from "react";
import "./Introduction.css";

class Introduction extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            metadata: false,
            spec: false,
            port: false
        };
        this.handleChange = this.handleChange.bind(this);

    }

    handleChange(event) {
        console.log(event.target.value)
        if (event.target.value == "metadata") {
            this.setState({ metadata: !this.state.metadata })
        }
        if (event.target.value == "spec-selector") {
            this.setState({ spec: !this.state.spec })
        }
        if (event.target.value == "spec-ports") {
            this.setState({ port: !this.state.port })
        }
    }

    handleSubmit(event) {
        event.preventDefault();
        console.log("Submitted")
        // fetch('http://0.0.0.0:8081/test', {
        //     method: 'POST',
        //     headers: {
        //         'Accept': 'application/json',
        //         'Content-Type': 'application/json',
        //     },
        //     body: JSON.stringify({
        //         apiVersion: this.state.apiVersion,
        //         kind: this.state.kind,
        //         metadata: this.state.metadata,
        //         path: this.state.path,
        //     })
        // })
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <label>
                    <h2>Please enter the username of the Gitlab repository:</h2>
                    <input type="text" name="username" onChange={this.handleChange} />
                </label>
                <label>
                    <h2>Please enter the path to the code base:</h2>
                    <input type="text" name="localFilePath" onChange={this.handleChange} />
                </label>
                <label>
                    <h2>Please enter the namespace for this deployment:</h2>
                    <input type="text" name="namespace" onChange={this.handleChange} />
                </label>
                <label>
                    <h2>Please enter the desired name of the application:</h2>
                    <input type="text" name="name" onChange={this.handleChange} />
                </label>
                <h5>The application will generate the argoCD config and the pipeline for you on the click of submit.</h5>
                <input type="submit" value="submit" />
            </form>
        )
    }
}

export default Introduction;