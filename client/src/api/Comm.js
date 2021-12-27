import axios from "axios";
import { UnAuthRole } from "roles";

const server_url = 'http://localhost:3000'

// TODO objects.assign(context, message.load)
const Comm = {
        get: async (context, message) => {
            console.log(Object.assign(context, message.load))
            return await axios.get(server_url + message.kind, {params: Object.assign(context, message.load)})
            .catch(error => console.log(error))
        },
        post: async (context, message) => {
            console.log(JSON.stringify(Object.assign(context, message.load)))
            return await axios({
                method: 'post',
                url: server_url + message.kind,
                data: JSON.stringify(Object.assign(context, message.load))
            }).catch(error => console.log(error))
        }
}

const FContext = {
    email: 'no email',
    username: 'no username',
    group: -1,
    course: -1,
    role: UnAuthRole
}

//Todo MessageKind, MessageLoad as interfaces in Typescript
export { Comm, FContext /*, MessageKind*/};
