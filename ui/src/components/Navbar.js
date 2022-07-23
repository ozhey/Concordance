import '../styles/Navbar.css';
import logo from '../assets/Concordance.png';

function Navbar() {
    return (
        <div className="navbar">
            <img src={logo} alt="Logo" />
            <div className="navbar__item">Articles</div>
            <div className="navbar__item">Index</div>
            <div className="navbar__item">Create</div>
        </div>
    );
}

export default Navbar;