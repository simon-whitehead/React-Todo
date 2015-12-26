import uuid from 'node-uuid';
import React from 'react';
import NoteList from './NoteList.jsx';

export default class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            notes: () => {
                let notes = [];
                for (let i = 0; i < 30; i++)
                    notes.push({ id: uuid.v4(), task: 'Task #' + i.toString() });

                return notes;
            }()
        }
    }

    editNote = (id, task) => {
        const notes = this.state.notes.map((note) => {
            if (note.id === id)
            note.task = task;

        return note;
        });

        this.setState({notes});
    }

    deleteNote = (id) => {
        this.setState({
            notes: this.state.notes.filter((note) => note.id !== id)
        });
    }

    render() {
        const {notes} = this.state;

        return (
            <div>
                <NoteList 
                items={notes} 
                onNoteEdit={this.editNote} 
                onNoteDelete={this.deleteNote} />
            </div>
       );
    }
}
