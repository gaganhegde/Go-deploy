import React from "react";
import "./Slider.css";
import Introduction from "./Introduction";
import {
    Link,
    BrowserRouter as Router
} from "react-router-dom";



class Slider extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            currentindex: 1,
            viewedAll: false,
        };
        this.handleNext = this.handleNext.bind(this);
        this.handlePrevious = this.handlePrevious.bind(this);
    }


    handleNext() {
        let newIndex = this.state.currentindex + 1

        if (newIndex > 3) {

            this.setState({ currentindex: 1, viewedAll: true })
            console.log("was here-1")
        }
        else {
            this.setState({ currentindex: newIndex })
        }
    }
    handlePrevious() {
        let newIndex = this.state.currentindex - 1
        if (newIndex < 1) {
            this.setState({ currentindex: 3 })
        }
        else {
            this.setState({ currentindex: this.state.currentindex - 1 })
        }
    }
    render() {
        return (
            <div className="main">
                <div className="Slider">
                    {this.state.currentindex === 1 ? <div className="Slide"><h1>A little bit about the application:</h1>
                        <ol> <li>This tool uses GitOps which is a new standard to perform deployments on the cluster. </li><li>This tool acts as a starter kit to monitor and make deployments on the cluster creating a starter pipeline and manifest folder for you.</li><li>The manifests can be mofidied at anytime to suite your needs by adding additional service and deployments through the UI.</li><li>Just add the manifests to you Gitlab code repo and it should trigger a simple pipeline for you.</li></ol><h3>Discalimer:</h3>The code currently works only for Go deployments.</div> : <div className="Slide-none"></div>}
                    {this.state.currentindex === 2 ? <div className="Slide"><h1>Before we start we need a few things to be kept handy:</h1><ul><li>To begin please make sure you have a GitLab repository with your code you plan to deploy available?</li><li>Please keep available the name you desire to call your deployment and the choice of port desired.</li><li>You can add services and dployment files to the configuration directory at a later point too.</li> <li>For security purposes it is reccommended that you create a github access token to auhtorise the pipelines to build and push images to the Repo.</li></ul> </div> : <div className="Slide-none"></div>}
                    {this.state.currentindex === 3 ? <div className="Slide"><h1>Once Created kindly copy the files into the  code repo and create a pull request</h1><h3>You  are now Ready to deploy using argoCD. The application will create you deployment.yaml, the  service.yaml, the ingress.yaml, addtional namespaces and the argoCD CRD</h3></div> : <div className="Slide-none"></div>}
                </div>
                <div>
                    <button onClick={this.handlePrevious}>Back</button>
                    <button onClick={this.handleNext}>Next</button>
                </div>
                <div>{this.state.viewedAll ? <Router><Link component={Introduction}>Continue</Link></Router> : <div className="Slide-none"></div>}</div>
            </div>)
    }
}

export default Slider;