
import {BrowserRouter, Route, Routes} from "react-router-dom";
import About from "./pages/About";
import axios from "axios";
import Comm from "./Comm";

function App() {

    async function fetch(){
        const response = await Comm.getSth()
        console.log(response)
    }

    return (
        <div className="App">
            <button onClick={fetch}>get</button>
        </div>
        // <BrowserRouter>
        //     <Routes>
        //         <Route path="/about" element={<About/>}/>
        //     </Routes>
        // </BrowserRouter>
    );
}

export default App;
