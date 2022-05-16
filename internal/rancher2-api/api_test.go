/*
 * Copyright 2022 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rancher2_api

import (
	"analytics-flow-engine/internal/lib"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"testing"
	"time"
)

func TestRancher2_createPersistentVolumeClaim(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Print("Error loading .env file")
	}
	driver := NewRancher2(
		lib.GetEnv("RANCHER2_ENDPOINT", ""),
		lib.GetEnv("RANCHER2_ACCESS_KEY", ""),
		lib.GetEnv("RANCHER2_SECRET_KEY", ""),
		lib.GetEnv("RANCHER2_STACK_ID", ""),
		lib.GetEnv("ZOOKEEPER", ""),
	)
	name := "test"
	err = driver.createPersistentVolumeClaim(name)
	if err != nil {
		fmt.Println(err.Error())
	}
	time.Sleep(3 * time.Second)

	err = driver.deletePersistentVolumeClaim(name)
	if err != nil {
		fmt.Println(err.Error())
	}

}
