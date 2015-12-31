import React from 'react';

import uuid from 'node-uuid';

import TaskList from './TaskList';
import TaskListActions from '../actions/TaskListActions';
import TaskListStore from '../stores/TaskListStore';

export default class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = TaskListStore.getState();
    }

    componentDidMount() {
        TaskListStore.listen(this.listStoreChanged);
    }

    componentWillUnmount() {
        TaskListStore.unlisten(this.listStoreChanged);
    }

    listStoreChanged = (state) => {
        this.setState(state);
    }

    addList = () => {
        TaskListActions.create({name: "New List"});
    }

    editListName = (id, name) => {
        TaskListActions.update({id, name});
    }

    deleteList = (id) => {
        TaskListActions.delete(id);
    }

    render() {
        const {lists} = this.state;

        return (
            <div>
                <button className="add-list" onClick={this.addList}>+ Add List</button>
                <ul className="list-container">
                    {lists.map(this.renderList, this)}
                </ul>
                <a href="/logout">Log out</a>
            </div>
       );
    }

    renderList(list, i) {
        return (
            <li key={i}> 
                <TaskList list={list} key={list.id} onEditListName={this.editListName} />
            </li>
        );
    }
}
