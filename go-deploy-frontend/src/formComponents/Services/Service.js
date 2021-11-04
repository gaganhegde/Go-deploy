import React from "react";
import Metadata from "../metadata";
import SpecSelector from "./Spec";
import Ports from "./ports";
import "./Service.css";

class Service extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            metadata: false,
            spec: false,
            port: false
        };
        this.handleChange = this.handleChange.bind(this);
        this.metavalue = this.metavalue.bind(this);

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
    metavalue(name) {
        console.log(name)
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
                    <Metadata metavalue={this.metavalue} />
                </div>
                <label><h2>Spec-selector:</h2><input value="spec-ports" type="checkbox" checked={this.state.metadata} onChange={this.handleChange}></input></label>
                <div style={{ display: this.state.port ? "block" : "none" }}>
                    <SpecSelector />
                </div>
                <label><h2>Spec-ports:</h2><input value="spec-selector" type="checkbox" checked={this.state.metadata} onChange={this.handleChange}></input></label>
                <div style={{ display: this.state.spec ? "block" : "none" }}>
                    <Ports />
                </div>
                <input type="submit" value="submit" />
            </form>
        )
    }
}

export default Service;