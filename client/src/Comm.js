import axios from "axios";

export default class Comm{
    static async getSth(){
        const response = await axios.get('http://localhost:3000/v1/admin/get/algo')

        return response
    }
}