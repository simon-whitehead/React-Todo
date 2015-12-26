import uuid from 'node-uuid';
import React from 'react';
import NoteList from './NoteList.jsx';

const notes = [
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
];

export default class App extends React.Component {
  render() {
    return (
      <div>
        <NoteList items={notes} />
      </div>
    );
  }
}
