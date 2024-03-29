import { Link } from "react-router-dom";
import { useContext } from "react";

import { AuthContext } from "../contexts/AuthContext";

import Logo from "../assets/logo.png";

const Navbar = () => {
  const { isLogin } = useContext(AuthContext);

  console.log(isLogin);

  return (
    <nav className="bg-gray-700/25 flex justify-between items-center py-3 px-8 navbar">
      <Link to="/" className="text-3xl font-semibold text-white">
        <img src={Logo} style={{ width: "80%" }} />
      </Link>
      <div>
        <div className="flex gap-5">
          <Link to="/dashboard" className="text-xl font-medium text-white">
            <p>Account</p>
          </Link>

          {!isLogin && (
            <div className="flex gap-5">
              <Link to="/sign-in" className="text-xl font-medium text-white">
                <p>Sign In</p>
              </Link>
              <Link to="/register" className="text-xl font-medium text-white">
                <p>Register</p>
              </Link>
            </div>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
