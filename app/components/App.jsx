import uuid from 'node-uuid';
import React from 'react';
import NoteList from './NoteList.jsx';
import NoteActions from '../actions/NoteActions';
import NoteStore from '../stores/NoteStore';

export default class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = NoteStore.getState();
    }

    componentDidMount() {
        NoteStore.listen(this.storeChanged);
    }

    componentWillUnmount() {
        NoteStore.unlisten(this.storeChanged);
    }

    storeChanged = (state) => {
        this.setState(state);
    }

    editNote = (id, task) => {
        NoteActions.update({id, task});
    }

    deleteNote = (id) => {
        NoteActions.delete(id);
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
