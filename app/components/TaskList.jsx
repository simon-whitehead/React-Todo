import React from "react";

import uuid from 'node-uuid';

import Editable from './Editable';

import Task from "./Task.jsx";
import TaskActions from "../actions/TaskActions";

import TaskStore from "../stores/TaskStore";
import TaskListStore from "../stores/TaskListStore";

export default class TaskList extends React.Component {
    constructor(props) {
        super(props);

        this.state = TaskStore.getState();
    }

    componentDidMount() {
        TaskStore.listen(this.taskStoreChanged);
    }

    componentWillUnmount() {
        TaskStore.unlisten(this.taskStoreChanged);
    }

    taskStoreChanged = (state) => {
        this.setState(state);
    }

    addTask = (list, task) => {
        TaskActions.create({list, task});
    }

    editTask = (id, task) => {
        TaskActions.update({id, task});
    }

    deleteTask = (id) => {
        TaskActions.delete(id);
    }

    render() {
	const list = this.props.list;
        const tasks = this.state.tasks.filter((t) => t.list_id === list.id);

	return (
            <div>
                <Editable value={list.name} onEditCompleted={this.props.onEditListName.bind(null, list.id)} />
                <button className="add-task" onClick={this.addTask.bind(null, list, "New Task")}>+ Add Task</button>
                <ul className="tasks">
                    {tasks.map(this.renderNote, this)}
                </ul>
            </div>
	);
    }

    renderNote(task, i) {
	return (
		<li className="task" key={i}>
			<Task 
                        task={task.task} 
                        onEditTask={this.editTask.bind(null, task.id)} 
                        onDeleteTask={this.deleteTask.bind(null, task.id)} key={task.id} />
		</li>
	);
    }
}
