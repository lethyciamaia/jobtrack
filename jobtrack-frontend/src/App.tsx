import React from 'react';
import logo from './logo.svg';
import './App.css';
import Header from './components/Header/header';
import MainView from './components/MainView/mainView';

function App() {
  return (
    <div className="App">
      <Header />
      <MainView />
    </div>
  );
}

export default App;
