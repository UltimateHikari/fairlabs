import axios from "axios";

export default class Comm{
    static async getAlgo(){
        return (await axios.get('http://localhost:3000/v1/admin/get/algo')).data
    }

    static async getAllCourses(){

    }

    static async getMyCourses(){

    }

    static async getProgress(){

    }

    static async getGoals(){

    }



}