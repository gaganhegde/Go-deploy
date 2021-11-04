import React from "react";
import Metadata from "../metadata";
import Spec from "./spec";

class Deployment extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            metadata: false
        };
        this.handleChange = this.handleChange.bind(this);
        this.metavalue = this.metavalue.bind(this);


    }

    handleChange(event) {
        if (event.target.value == "metadata") {
            this.setState({ metadata: !this.state.metadata })
            console.log("was here")
        }
    }
    metavalue(name) {
        console.log(name)
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
                    <h2>apiVersion:</h2>
                    <input type="text" name="apiVersion" onChange={this.handleChange} />
                </label>
                <label>
                    <h2>kind:</h2>
                    <input type="text" name="kind" onChange={this.handleChange} />
                </label>
                <label><h2>Metadata:</h2><input value="metadata" type="checkbox" checked={this.state.metadata} onChange={this.handleChange}></input></label>
                <div style={{ display: this.state.metadata ? "block" : "none" }}>
                    <Metadata onChange={this.metavalue} />
                </div>
                <Spec />
                <input type="submit" value="submit" />
            </form>
        )
    }
}

export default Deployment;