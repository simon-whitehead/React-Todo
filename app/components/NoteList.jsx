import React from "react";
import Note from "./Note.jsx";

export default class NoteList extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
	const notes = this.props.items;

	return (
    	    <ul className="notes">
	        {notes.map(this.renderNote, this)}
	    </ul>
	);
    }

    renderNote(note) {
	return (
		<li className="note" key={note.id}>
			<Note 
                        task={note.task} 
                        onEdit={this.props.onNoteEdit.bind(null, note.id)} 
                        onDelete={this.props.onNoteDelete.bind(null, note.id)} />
		</li>
	);
    }
}
