import uuid from 'node-uuid';
import React from 'react';
import TaskList from './TaskList.jsx';
import TaskActions from '../actions/TaskActions';
import TaskStore from '../stores/TaskStore';

export default class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = TaskStore.getState();
    }

    componentDidMount() {
        TaskStore.listen(this.storeChanged);
    }

    componentWillUnmount() {
        TaskStore.unlisten(this.storeChanged);
    }

    storeChanged = (state) => {
        this.setState(state);
    }

    editTask = (id, task) => {
        TaskActions.update({id, task});
    }

    deleteTask = (id) => {
        TaskActions.delete(id);
    }

    render() {
        const {tasks} = this.state;

        return (
            <div>
                <TaskList 
                items={tasks} 
                onEditTask={this.editTask} 
                onDeleteTask={this.deleteTask} />
            </div>
       );
    }
}
