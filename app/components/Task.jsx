import React from 'react';

export default class Task extends React.Component {
    constructor(props) {
	super(props);

	this.state = {
	    editing: false
	};
    }
    render() {
	if(this.state.editing) {
	    return this.renderEdit();
	}

	return this.renderTask();
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
    edit = () => {
	this.setState({
	    editing: true
	});
    }
    checkEnter = (e) => {
	if(e.key === 'Enter') {
	    this.finishEdit(e);
	}
    }
    finishEdit = (e) => {
	this.props.onEditTask(e.target.value);

	this.setState({
	    editing: false
	});
    }
}
