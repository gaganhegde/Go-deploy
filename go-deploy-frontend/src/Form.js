import React from 'react';
import "./Form.css";
import Deployment from "./formComponents/Deployments/deployment";

class Form extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            apiVersion: "",
            kind: "",
            image: "",
            ports: "",
            path: "",
            metadataName: false,
            metadatLabel: false,
        }
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    handleChange(event) {
        console.log(event.target.value)
        if (event.target.name == "metadata") {
            this.setState({
                metadataName: true,
            })
        }
    }

    handleSubmit(event) {
        event.preventDefault();
        fetch('http://0.0.0.0:8081/test', {
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                apiVersion: this.state.apiVersion,
                kind: this.state.kind,
                metadata: this.state.metadata,
                path: this.state.path,
            })
        })
        console.log(JSON.stringify({
            apiVersion: this.state.apiVersion,
            kind: this.state.kind,
            metadata: this.state.metadataName,
            path: this.state.path,
        }))
    }

    render() {
        const metaname = this.state.metadataName ? <div>
            imageName :
            <input type="text" value={this.state.name} onChange={this.handleChange} />
            <h3>labels:</h3>
            <input type="checkbox" name="labels" onChange={this.handleChange} />
        </div> : null;
        const metalabel = this.state.metadataLabel ? <div>
            app :
        <input type="text" value={this.state.name} onChange={this.handleChange} />
        role :
        <input type="text" value={this.state.name} onChange={this.handleChange} />
        tier :
        <input type="text" value={this.state.name} onChange={this.handleChange} />
        </div> : null;


        return <div >
            <h1>Create a deployment</h1>
            <form onSubmit={this.handleSubmit}>
                <Deployment />
                <input type="submit" value="submit" />
            </form>
        </div>
    }
}

export default Form;