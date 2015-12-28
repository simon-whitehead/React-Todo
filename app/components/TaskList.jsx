import React from "react";
import Task from "./Task.jsx";

export default class TaskList extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
	const tasks = this.props.items;

	return (
            <div>
                <button className="add-task" onClick={this.props.onAddTask.bind(null, "New Task")}>+</button>
                <ul className="tasks">
                    {tasks.map(this.renderNote, this)}
                </ul>
            </div>
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
