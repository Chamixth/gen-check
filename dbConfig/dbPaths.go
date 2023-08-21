package dbConfig

import (
	mongo "go.mongodb.org/mongo-driver/mongo"
)

var DATABASE *mongo.Database

// const DATABASE_URL = "mongodb+srv://FiberUser:isk19940710@cluster0.yxgn5.mongodb.net/?retryWrites=true&w=majority"
//const DATABASE_URL = "mongodb+srv://AnujanR:4rW0ZaiPmBtkbrXA@evoar.qcmphjl.mongodb.net/?retryWrites=true&w=majority"
const DATABASE_URL = "mongodb+srv://chamith:123@cluster0.ujlq82i.mongodb.net/?retryWrites=true&w=majority"

const DATABASE_NAME ="chaithApp"
