const QueueMessage = () => {
    return {
        kind: '/v1/tasks/queue',
        load: {}
    }
}
const MyCoursesMessage = () => {
    return {
        kind: '/v1/tasks/mycourses',
        load: {}
    }
}
const AllCoursesMessage = () => {
    return {
        kind: '/v1/tasks/allcourses',
        load: {}
    }
}
const FollowMessage = (course_id) => {
    return {
        kind: '/v1/tasks/follow',
        load: {
            id: course_id
        }
    }
}
const SubmitMessage = (intent, tasks) => {
    return {
        kind: '/v1/tasks/submit',
        load: {
            intent: intent,
            tasks: tasks
        }
    }
}
const QueryMessage = (intent) => {
    return {
        kind: '/v1/tasks/query',
        load: {
            intent: intent
        }
    }
}

export {QueueMessage, MyCoursesMessage, AllCoursesMessage, FollowMessage, SubmitMessage, QueryMessage}