#!/bin/bash

echo $(tr -d '\n' < $1) > $1
