#!/bin/bash

set -eu

version="${1:?missing version}"
git tag "$(cat "$version")"
