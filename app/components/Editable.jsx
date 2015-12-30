import React from 'react';

export default class Editable extends React.Component {
    constructor(props) {
	super(props);

	this.state = {
	    editing: false,
            value: props.value ? props.value : ''
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
	defaultValue={this.state.value}
	onBlur={this.finishEdit}
	onKeyPress={this.checkEnter} />;
    }
    renderTask = () => {
	return (
            <div onClick={this.edit}>
                <span className="editable-static">{this.state.value}</span>
            </div>
       );
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
	this.props.onEditCompleted(e.target.value);

	this.setState({
	    editing: false,
            value: e.target.value
	});
    }
}
