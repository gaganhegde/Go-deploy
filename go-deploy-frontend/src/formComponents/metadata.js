import React from "react";
import Labels from "./Deployments/labels";

class Metadata extends React.Component {
    constructor(props) {
        super(props);
        this.onchange = this.onchange.bind(this);
    }
    onchange(event) {
        const Name = event.target.value;
        this.props.metavalue(Name);
        console.log("was here -2")
    }

    render() {
        return (
            <div>
                <label>
                    <h3>Name:</h3>
                    <input type="text" name="apiVersion" onChange={this.onchange} />
                </label>
                <label><h3>Labels:</h3></label>
                <Labels />
            </div>
        )
    }
}

export default Metadata;