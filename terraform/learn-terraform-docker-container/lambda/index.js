module.exports = {
  test: async (event) => {
    console.log("I have been called")
    return {
      statusCode: 200,
      body: JSON.stringify(
        {
          message: "Go Serverless v3.0! Your function executed successfully!",
          input: event,
        },
        null,
        2
      ),
    }
  },
  run: async (event) => {
  console.log("I have been called")
  return {
    statusCode: 200,
    body: JSON.stringify(
      {
        message: "Go Serverless v3.0! Your function executed successfully!",
        input: event,
      },
      null,
      2
    ),
  }
}}

https://learn.hashicorp.com/tutorials/terraform/aws-change