#bin/bash
version=v1.0.3
git tag -a ${version} -m '${version}'

git push origin ${version}
