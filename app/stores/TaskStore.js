import uuid from 'node-uuid';
import alt from '../libs/alt';
import TaskActions from '../actions/TaskActions';

var $ = require('jquery');

class TaskStore {
    constructor() {
        this.bindActions(TaskActions);

        this.tasks = [];
    }
    create(task) {
        const tasks = this.tasks;

        task.id = uuid.v4();

        this.setState({
            tasks: tasks.concat(task)
        });

        $.post('/api/create-task', task);
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
}

export default alt.createStore(TaskStore, 'TaskStore');
