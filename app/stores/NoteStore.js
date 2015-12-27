import uuid from 'node-uuid';
import alt from '../libs/alt';
import NoteActions from '../actions/NoteActions';

class NoteStore {
    constructor() {
        this.bindActions(NoteActions);

        this.notes = () => {
            let notes = [];
            for (let i = 0; i < 30; i++)
                notes.push({ id: uuid.v4(), task: 'Task #' + i.toString() });

            return notes;
        }();
    }
    create(note) {
        const notes = this.notes;

        note.id = uuid.v4();

        this.setState({
            notes: notes.concat(note)
        });
    }
    update({id, task}) {
        const notes = this.notes.map((note) => {
            if(note.id === id) {
                note.task = task;
            }

            return note;
        });

        this.setState({notes});
    }
    delete(id) {
        this.setState({
            notes: this.notes.filter((note) => note.id !== id)
        });
    }
}

export default alt.createStore(NoteStore, 'NoteStore');
