/*
Copyright © 2023 Raymond onepiece010938@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/onepiece010938/go-line-message-analyzer/aws"
	"github.com/onepiece010938/go-line-message-analyzer/cmd"
)

var DLAMBDA *aws.DeployLambda

func main() {
	deploy := os.Getenv("DEPLOY_PLATFORM")
	if deploy == "lambda" {
		rootCtx, _ := context.WithCancel(context.Background()) //nolint
		DLAMBDA = aws.NewDeployLambda(rootCtx)
		lambda.Start(DLAMBDA.Handler)
	} else {
		cmd.Execute()
	}

}
