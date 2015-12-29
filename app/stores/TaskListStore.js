import uuid from 'node-uuid';
import alt from '../libs/alt';
import TaskListActions from '../actions/TaskListActions';

var $ = require('jquery');

class TaskListStore {
    constructor() {
        this.bindActions(TaskListActions);

        this.lists = [];
    }
    create(list) {
        const lists = this.lists;

        list.id = uuid.v4();

        this.setState({
            lists: lists.concat(list)
        });

        // Persist the list here
    }
    update({id, name}) {
        const lists = this.lists.map((l) => {
            if (l.id === id) {
                l.name = name;
            }

            return l;
        });

        this.setState({lists});
    }
    delete(id) {
        this.setState({
            lists: this.lists.filter((l) => l.id !== id)
        });
    }
}

export default alt.createStore(TaskListStore, 'TaskListStore');
