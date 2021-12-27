import axios from "axios";

const server_url = 'http://localhost:3000'

// TODO objects.assign(context, message.load)
const Comm = {
        get: async (context, message) => {
            console.log(context)
            await axios({
                method: 'get',
                url: server_url + message.kind,
                body: JSON.stringify(Object.assign(context, message.load))
            }).then(response => response.data)
            .catch(error => console.log(error))
        },
        post: async (context, message) => {
            await axios({
                method: 'post',
                url: server_url + message.kind,
                body: JSON.stringify(Object.assign(context, message.load))
            }).then(response => console.log(response))
            .catch(error => console.log(error))
        }
}

// roles
// 1 == student
// 2 == teacher
// 3 == admin

const FContext = {
    email: 'no email',
    username: 'no username',
    group: -1,
    course: -1,
    role: -1
}

//Todo MessageKind, MessageLoad as interfaces in Typescript
export { Comm, FContext /*, MessageKind*/};
