# GeoQuery

Peta Pedia Library Package

## ENV Setting
![image](https://github.com/petapedia/geoquery/assets/11188109/1009c159-7cb2-4d0f-bf78-5ad590a307ad)

## Release
```sh
go get -u all
go mod tidy
git tag                                 #check current version
git tag v0.0.28                          #set tag version
git push origin --tags                  #push tag version to repo
go list -m github.com/petapedia/geoquery@v0.0.28   #publish to pkg dev, replace ORG/URL with your repo URL
```
