#bin/bash
version=1.0.3
git tag -a ${version} -m '${version}'

git push origin ${version}
