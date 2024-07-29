# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

setup() {
    load 'test_helper/bats-support/load'
    load 'test_helper/bats-assert/load'
    export NO_COLOR=1
    export NUV_BRANCH="0.1.0-testing"
    cd prereq
}

@test "download others" {
    rm -Rvf ~/.nuv/
    for o in linux darwin
    do for a in arm64 amd64
       do  
          run env OS=$o ARCH=$a ops
          assert_line "ensuring prerequisite 7zz" --partial
          assert test -e ~/.nuv/$o-$a/bin/7zz
       done
    done 
    run env OS=windows ops
    assert test -e ~/.nuv/windows-*/bin/7zz.exe
}

@test "ops prereq" {
    rm -Rvf ~/.nuv/
    run ops
    assert_line "ensuring prerequisite 7zz" --partial
    run ops info
    assert_line "7-˜Zip" --partial
}
