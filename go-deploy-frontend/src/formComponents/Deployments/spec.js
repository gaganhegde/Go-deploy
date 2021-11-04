import React from "react";
import ContainerElement from "./containers";

class Spec extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            containers: 0,
            fields: [],
        };
        this.handleChange = this.handleChange.bind(this);

    }
    handleChange(event) {
        const newArray = []
        for (let i = 0; i < event.target.value; i++) {
            newArray.push(<div><ContainerElement /></div>)
        }
        this.setState({ fields: newArray })
    }


    render() {
        return (
            <div>
                <label><h2>Containers:</h2>
                    <select onChange={this.handleChange}>
                        <option key="0" value={0}>Select number of containers:</option>
                        <option key="1" value={1}>1</option>
                        <option key="2" value={2}>2</option>
                        <option key="3" value={3}>3</option>
                    </select>
                </label>
                {this.state.fields}
            </div>
        )
    }
}

export default Spec;