import uuid from 'node-uuid';
import alt from '../libs/alt';
import TaskActions from '../actions/TaskActions';

class TaskStore {
    constructor() {
        this.bindActions(TaskActions);

        this.tasks = () => {
            let tasks = [];
            for (let i = 0; i < 30; i++)
                tasks.push({ id: uuid.v4(), task: 'Task #' + i.toString() });

            return tasks;
        }();
    }
    create(note) {
        const tasks = this.tasks;

        note.id = uuid.v4();

        this.setState({
            tasks: tasks.concat(note)
        });
    }
    update({id, task}) {
        const tasks = this.tasks.map((note) => {
            if(note.id === id) {
                note.task = task;
            }

            return note;
        });

        this.setState({tasks});
    }
    delete(id) {
        this.setState({
            tasks: this.tasks.filter((note) => note.id !== id)
        });
    }
}

export default alt.createStore(TaskStore, 'TaskStore');
