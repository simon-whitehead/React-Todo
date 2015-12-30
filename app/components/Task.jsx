import React from 'react';
import Editable from './Editable';

export default class Task extends React.Component {
    constructor(props) {
	super(props);

	this.state = {
	    editing: false
	};
    }
    render() {
        return (
            <div className="task-title">
                <Editable value={this.props.task} onEditCompleted={this.finishEdit.bind(null)} />
                <div className="delete-task">
                    {this.onDeleteTask ? this.renderDelete() : null}
                </div>
            </div>
        );
    }
    renderEdit = () => {
	return <input type="text"
	    autoFocus={true}
	defaultValue={this.props.task}
	onBlur={this.finishEdit}
	onKeyPress={this.checkEnter} />;
    }
    renderTask = () => {
	const onDelete = this.props.onDeleteTask;

	return (
		<div onClick={this.edit}>
		<span className="task">{this.props.task}</span>
		{onDelete ? this.renderDelete() : null }
		</div>
	       );
    }
    renderDelete = () => {
	return <button className="delete" onClick={this.props.onDeleteTask}>x</button>;
    }
    finishEdit = (value) => {
	this.props.onEditTask(value);

	this.setState({
	    editing: false
	});
    }
}
