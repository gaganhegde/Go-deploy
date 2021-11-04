import logo from './logo.svg';
import './App.css';
import Form from "./Form";
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';
import React from 'react';
import Deployment from "./formComponents/Deployments/deployment";
import Service from "./formComponents/Services/Service";
class App extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      value: 1
    }
    this.handleChange = this.handleChange.bind(this);
  }

  handleChange(event) {
    this.setState({
      value: event.target.value,
    })
  }
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <h1>Go-deploy</h1>
        </header>
        <body>
          <Tabs>
            <TabList>
              <Tab><div className="Tab"><h3>Deployment</h3></div></Tab>
              <Tab><div className="Tab"><h3>Service</h3></div></Tab>
              <Tab><div className="Tab"><h3>Pods/Patches</h3></div></Tab>
              <Tab><div className="Tab"><h3>Cluster Metrics</h3></div></Tab>
            </TabList>

            <TabPanel>
              <div className="Tabs"> <Deployment /></div>
            </TabPanel>
            <TabPanel>
              <div className="Tabs"><Service /></div>
            </TabPanel>

          </Tabs>
        </body>
      </div>
    );
  }
}

export default App;
