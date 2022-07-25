import './App.css';
import Navbar from "./components/Navbar";
import Articles from "./components/Articles";

function App() {
    return (
        <div className="app">
            <Navbar/>
            <div className="page-container">
                <Articles/>
            </div>
        </div>
    );
}

export default App;
