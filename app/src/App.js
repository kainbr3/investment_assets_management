import {
  Route, 
  Routes
} from "react-router-dom";

import Login from "./pages/login";
import Dashboard from "./pages/dashboard";

function App() {
  return (
    <>
      <Routes>
        <Route exact path="/" element={<Login />} />
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </>
  );
}

export default App;
