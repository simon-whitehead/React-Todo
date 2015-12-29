import uuid from 'node-uuid';
import alt from '../libs/alt';
import TaskActions from '../actions/TaskActions';

var $ = require('jquery');

class TaskStore {
    constructor() {
        this.bindActions(TaskActions);

        this.tasks = [];
    }
    create({list, task}) {
        const tasks = this.tasks;

        this.setState({
            tasks: tasks.concat({
                id: uuid.v4(),
                task: 'New task',
                list_id: list.id
            })
        });
    }
    update({id, task}) {
        const tasks = this.tasks.map((t) => {
            if(t.id === id) {
                t.task = task;
            }

            return t;
        });

        this.setState({tasks});
    }
    delete(id) {
        this.setState({
            tasks: this.tasks.filter((t) => t.id !== id)
        });
    }
    getTasksForList(list) {
        return this.tasks.filter((t) => t.list_id === list.id);    
    }
}

export default alt.createStore(TaskStore, 'TaskStore');
