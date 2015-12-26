import uuid from 'node-uuid';
import React from 'react';
import NoteList from './NoteList.jsx';

export default class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            notes: [
            {
                id: uuid.v4(),
                task: 'Learn Webpack'
            },
            {
                id: uuid.v4(),
                task: 'Learn React'
            },
            {
                id: uuid.v4(),
                task: 'Do laundry'
            }
            ]
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
