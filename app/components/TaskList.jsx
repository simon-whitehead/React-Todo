import React from "react";
import Task from "./Task.jsx";

export default class TaskList extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
	const tasks = this.props.items;

	return (
    	    <ul className="tasks">
	        {tasks.map(this.renderNote, this)}
	    </ul>
	);
    }

    renderNote(task) {
	return (
		<li className="task" key={task.id}>
			<Task 
                        task={task.task} 
                        onEditTask={this.props.onEditTask.bind(null, task.id)} 
                        onDeleteTask={this.props.onDeleteTask.bind(null, task.id)} />
		</li>
	);
    }
}
