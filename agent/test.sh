#!/bin/bash

echo -n $RANDOM

exit `expr $RANDOM % 3`
