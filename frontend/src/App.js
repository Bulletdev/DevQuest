import React from 'react';
import LessonList from './components/LessonList';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <h1>CodeQuest</h1>
      </header>
      <main>
        <LessonList />
      </main>
    </div>
  );
}

export default App;
