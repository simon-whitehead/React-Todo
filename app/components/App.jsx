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

    render() {
        return (
          <div>
            <NoteList items={notes} />
          </div>
        );
    }
}
