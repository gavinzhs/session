package main

//
//import (
//    "encoding/base64"
//    "encoding/json"
//    "fmt"
//    "github.com/go-martini/martini"
//    "github.com/qiniu/api/auth/digest"
//    qiniuio "github.com/qiniu/api/io"
//    "github.com/qiniu/api/rs"
//    "gopkg.in/mgo.v2"
//    "gopkg.in/mgo.v2/bson"
//    "io"
//    "io/ioutil"
//    "log"
//    "net/http"
//    "net/url"
//    "strings"
//)
//
//const (
//    ACCESS_KEY_OCCALL = "QwbwQWwnggA04VWQjo_mcep0Z_aS1VO07zmcN2HZ"
//    SECRET_KEY_OCCALL = "M5VTtldMM3kaARCQbvitOhO3xQr2Tp9H-PRzY6Jl"
//)
//
//var (
//    mac_occall = &digest.Mac{
//        AccessKey: ACCESS_KEY_OCCALL,
//        SecretKey: []byte(SECRET_KEY_OCCALL),
//    }
//
//    BUCKET_OCCALL string
//
//    CALLBACK_URL_OCCALL string
//    QINIU_DOMAIN_OCCALL string
//)
//
//func init() {
//    if martini.Env == martini.Prod {
//        BUCKET_OCCALL = "occall"
//        QINIU_DOMAIN_OCCALL = "7xk0mz.dl1.z0.glb.clouddn.com"
//        CALLBACK_URL_OCCALL = "http://occall.com/qiniu/callback"
//    } else {
//        BUCKET_OCCALL = "occall-test"
//        QINIU_DOMAIN_OCCALL = "7xk0n0.dl1.z0.glb.clouddn.com"
//        CALLBACK_URL_OCCALL = "http://beta.occall.com/qiniu/callback"
//    }
//}
//
///**
// * 生成上传Token
// */
//func getUploadTokenByOccall(web *Web) (int, string) {
//    policy := &rs.PutPolicy{
//        Scope: BUCKET_OCCALL,
//    }
//    uploadToken := policy.Token(mac_occall)
//
//    return web.Json(200, J{"token": uploadToken})
//}
//
///**
// * 生成转码Token
// */
//func getRequestTokenByOccall(req *http.Request) (string, error) {
//
//    token, err := mac_occall.SignRequest(req, true)
//    if err != nil {
//        return "", err
//    }
//
//    return token, err
//}
//
///**
// * 生成下载Url
// * domain：域名
// * key：文件key
// */
//func downloadUrlByOccall(domain, key string) string {
//    baseUrl := rs.MakeBaseUrl(domain, key)
//    policy := rs.GetPolicy{}
//    return policy.MakeRequest(baseUrl, nil)
//}
//
///**
// * 获取文件信息
// * bucket：空间名称
// * key：文件key
// */
//func getInfoByOccall(bucket string, key string) (rs.Entry, error) {
//    client := rs.New(mac_occall)
//
//    ret, err := client.Stat(nil, bucket, key)
//    if err != nil {
//        //产生错误
//        log.Println("rs.Stat failed:", err)
//        return ret, err
//    }
//    //处理返回值
//    log.Println(ret)
//
//    return ret, nil
//}
//
///**
// * 删除文件
// * bucket：空间名称
// * key：文件key
// */
//func deleteFileByOccall(bucket string, key string) error {
//    client := rs.New(mac_occall)
//
//    err := client.Delete(nil, bucket, key)
//
//    return err
//}
//
///**
// * 获取视频信息接口
// * GET <AvDownloadURI>?avinfo
// */
//func getVideoInfoByOccall(key string) (*VideoInfo, error) {
//    url := fmt.Sprintf("http://%s/%s?avinfo", QINIU_DOMAIN_OCCALL, key)
//    resp, err := http.Get(url)
//    if err != nil {
//        fmt.Printf("get video info error: %v\n", err)
//        return nil, err
//    }
//
//    defer resp.Body.Close()
//
//    data, err := ioutil.ReadAll(resp.Body)
//    if err != nil {
//        fmt.Printf("get video info response error: %v\n", err)
//        return nil, err
//    }
//
//    //fmt.Println("video info: ", string(data))
//
//    var info *VideoInfo
//    err = json.Unmarshal(data, &info)
//    if err != nil {
//        fmt.Printf("get video info unmashal json error: %v\n", err)
//        return nil, err
//    }
//
//    return info, nil
//}
//
///**
// * 转码
// *
// * 在数据处理命令后用管道符 | 拼接 saveas/<encodedEntryURI> 指令，指示七牛服务器使用 EntryURI 格式中指定的 Bucket 与 Key 来保存处理结果。
// * 例如avthumb/flv|saveas/cWJ1Y2tldDpxa2V5，是将上传的视频文件转码成flv格式后存储为qbucket:qkey，
// * 其中cWJ1Y2tldDpxa2V5是qbucket:qkey的UrlSafe-Base64编码结果。
// */
//func encodeByOccall(key string) (string, error) {
//
//    info, err := getVideoInfo(key)
//    if err != nil {
//        fmt.Printf("get video info error: %v\n", err)
//        return "", fmt.Errorf("get video info error: %v", err)
//    }
//    if info.Error != "" {
//        fmt.Println("get video info response error: %v\n", info.Error)
//    }
//
//    width := 480
//    height := 320
//    for _, s := range info.Streams {
//        if s.CodecType == CODEC_TYPE_VIDEO {
//            width = s.Width
//            height = s.Height
//        }
//    }
//
//    target := fmt.Sprintf("%v:%v", BUCKET_OCCALL, "mp4_"+key)
//    targetKey := base64.URLEncoding.EncodeToString([]byte(target))
//
//    imgTarget := fmt.Sprintf("%v:%v", BUCKET_OCCALL, "img_"+key)
//    imgTargetKey := base64.URLEncoding.EncodeToString([]byte(imgTarget))
//
//    //fmt.Println("target key:", targetKey)
//    q := url.Values{
//        "bucket":    {BUCKET_OCCALL},
//        "key":       {key},
//        "fops":      {fmt.Sprintf("avthumb/mp4|saveas/%v;vframe/jpg/offset/1/w/%v/h/%v|saveas/%v", targetKey, width, height, imgTargetKey)},
//        "notifyURL": {"http://ip/qiniu/callback"},
//    }
//
//    host := "api.qiniu.com"
//    api := "pfop"
//
//    url := fmt.Sprintf("http://%v/%v", host, api)
//    req, err := http.NewRequest("POST", url, strings.NewReader(q.Encode()))
//    if err != nil {
//        return "", fmt.Errorf("request err %s, %s\n", url, err)
//    }
//
//    token, err := mac_occall.SignRequest(req, true)
//    if err != nil {
//        return "", fmt.Errorf("generate request token error: %v", err)
//    }
//    req.Header.Set("Host", "api.qiniu.com")
//    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//    req.Header.Add("Authorization", fmt.Sprintf("QBox %v", token))
//
//    resp, err := http.DefaultClient.Do(req)
//    defer resp.Body.Close()
//
//    if resp.StatusCode != 200 {
//        return "", fmt.Errorf("[http] status err %s, %d\n", url, resp.StatusCode)
//    }
//    b, err := ioutil.ReadAll(resp.Body)
//    if err != nil {
//        return "", fmt.Errorf("[http] read err %s, %s\n", url, err)
//    }
//
//    var res *FopsRes
//    err = json.Unmarshal(b, &res)
//    if err != nil {
//        fmt.Printf("unmarshal response error: %v\n", err)
//        fmt.Printf("response: %v\n", string(b))
//        return "", fmt.Errorf("unmarshal response error: %v", err)
//    }
//    return res.PersistentId, nil
//}
//
//func getPersistentStatByByOccall(persistentId string) (string, error) {
//
//    host := "api.qiniu.com"
//    api := "/status/get/prefop"
//    url := fmt.Sprintf("http://%v%v?id=%v", host, api, persistentId)
//
//    http.Get(url)
//    req, err := http.NewRequest("GET", url, nil)
//    if err != nil {
//        return "", fmt.Errorf("request err %s, %s\n", url, err)
//    }
//
//    //	req.Header.Set("Host", "api.qiniu.com")
//
//    resp, err := http.DefaultClient.Do(req)
//    defer resp.Body.Close()
//
//    if resp.StatusCode != 200 {
//        return "", fmt.Errorf("[http] status err %s, %d\n", url, resp.StatusCode)
//    }
//    b, err := ioutil.ReadAll(resp.Body)
//    if err != nil {
//        return "", fmt.Errorf("[http] read err %s, %s\n", url, err)
//    }
//
//    return string(b), nil
//}
//
//func getPersistentResultByOccall(fop string) (string, error) {
//    host := "api.qiniu.com"
//    api := "/p/1"
//    url := fmt.Sprintf("http://%v%v/%v", host, api, fop)
//
//    http.Get(url)
//    req, err := http.NewRequest("GET", url, nil)
//    if err != nil {
//        return "", fmt.Errorf("request err %s, %s\n", url, err)
//    }
//
//    req.Header.Set("Host", "api.qiniu.com")
//
//    resp, err := http.DefaultClient.Do(req)
//    defer resp.Body.Close()
//
//    if resp.StatusCode != 200 {
//        return "", fmt.Errorf("[http] status err %s, %d\n", url, resp.StatusCode)
//    }
//    b, err := ioutil.ReadAll(resp.Body)
//    if err != nil {
//        return "", fmt.Errorf("[http] read err %s, %s\n", url, err)
//    }
//
//    return string(b), nil
//}
//
//func uploadByOccall() {
//
//    policy := &rs.PutPolicy{
//        Scope:               BUCKET_OCCALL,
//        CallbackUrl:         CALLBACK_URL_OCCALL,
//        CallbackBody:        "key=$(key)&hash=$(etag)",
//        PersistentOps:       "avthumb/m3u8/segment/10",
//        PersistentNotifyUrl: CALLBACK_URL_OCCALL,
//    }
//    uploadToken := policy.Token(mac_occall)
//    fmt.Println("upload token:", uploadToken)
//
//    var err error
//    var ret qiniuio.PutRet
//    var extra = &qiniuio.PutExtra{
//        //Params:    params,
//        //MimeType:  mieType,
//        //Crc32:     crc32,
//        //CheckCrc:  CheckCrc,
//    }
//
//    // ret       变量用于存取返回的信息，详情见 io.PutRet
//    // uptoken   为业务服务器端生成的上传口令
//    // key       为文件存储的标识
//    // r         为io.Reader类型，用于从其读取数据
//    // extra     为上传文件的额外信息,可为空， 详情见 io.PutExtra, 可选
//
//    file := "2.mp4"
//    err = qiniuio.PutFile(nil, &ret, uploadToken, fmt.Sprintf("test-%v", tick()), file, extra)
//
//    if err != nil {
//        //上传产生错误
//        log.Println("io.Put failed:", err)
//        return
//    }
//
//    //上传成功，处理返回值
//    log.Println(ret.Hash, ret.Key)
//}
//
//func callbackByOccallHandler(vo CallBack, web *Web, ds *Ds) (int, string) {
//
//    b, err := json.Marshal(vo)
//    chk(err)
//    fmt.Println("callback: ", string(b))
//
//    if vo.Code == 0 {
//        r, err := findResourceByKey(ds.se, vo.InputKey)
//        if err != nil {
//            if err == mgo.ErrNotFound {
//                fmt.Printf("encode callback error, resource not found: %v\n", vo.InputKey)
//                return web.Json(200, J{"success": true})
//            }
//            fmt.Printf("encode callback, find resource error: %v\n", err)
//            return web.Json(200, J{"success": true})
//        }
//
//        for _, item := range vo.Items {
//            if item.Code != 0 {
//                fmt.Printf("encode callback, item code error: %v\n", item.Code)
//                return web.Json(200, J{"success": true})
//            }
//            if strings.Index(item.Cmd, "avthumb/mp4") == 0 {
//                if item.Key != fmt.Sprintf("mp4_%v", vo.InputKey) {
//                    fmt.Printf("encode callback, video target key error: %v\n", vo.Items[0].Key)
//                    return web.Json(200, J{"success": true})
//                }
//            } else {
//                if item.Key != fmt.Sprintf("img_%v", vo.InputKey) {
//                    fmt.Printf("encode callback, img target key error: %v\n", vo.Items[0].Key)
//                    return web.Json(200, J{"success": true})
//                }
//            }
//        }
//
//        r.Mp4 = fmt.Sprintf("mp4_%v", vo.InputKey)
//        r.Poster = fmt.Sprintf("img_%v", vo.InputKey)
//        r.Status = RESOURCE_STATUS_VALID
//        err = updateResource(ds.se, r)
//        if err != nil {
//            fmt.Printf("encode callback, update resource status error: %v\n", err)
//            return web.Json(200, J{"success": true})
//        }
//    }
//
//    return web.Json(200, J{"success": true})
//}
//
//func uploadImgQiniuByOccallHandler(r *http.Request, ds *Ds, web *Web) (int, string) {
//    file, header, err := r.FormFile("bin")
//    chk(err)
//    defer file.Close()
//
//    resource, code, info := qiniuUpFileByOccall(ds, file, header.Filename, RESOURCE_TYPE_PIC)
//    if resource == nil {
//        return code, info
//    }
//
//    url := qiniuImgUrlByOccall(resource.Id.Hex())
//    j := J{
//        "id":       resource.Id.Hex(),
//        "name":     resource.Name,
//        "url":      url,
//        "download": url + "?attname=" + resource.Name,
//    }
//
//    return web.Json(200, j)
//}
//
//func uploadFileByOccallHandler(r *http.Request, ds *Ds, web *Web) (int, string) {
//    file, header, err := r.FormFile("file")
//    chk(err)
//    defer file.Close()
//
//    resource, code, info := qiniuUpFileByOccall(ds, file, header.Filename, RESOURCE_TYPE_FILE)
//    if resource == nil {
//        return code, info
//    }
//
//    url := fileDownloadByOccall(resource)
//    j := J{
//        "id":       resource.Id.Hex(),
//        "name":     resource.Name,
//        "url":      url,
//        "download": url + "?attname=" + resource.Name,
//    }
//
//    return web.Json(200, j)
//}
//
//func qiniuUpFileByOccall(ds *Ds, data io.Reader, filename string, tp int) (*Resource, int, string) {
//    // ret       变量用于存取返回的信息，详情见 io.PutRet
//    // uptoken   为业务服务器端生成的上传口令
//    // key       为文件存储的标识
//    // r         为io.Reader类型，用于从其读取数据
//    // extra     为上传文件的额外信息,可为空， 详情见 io.PutExtra, 可选
//
//    policy := &rs.PutPolicy{
//        Scope: BUCKET_OCCALL,
//    }
//
//    if tp == RESOURCE_TYPE_VIDEO {
//        policy.CallbackUrl = CALLBACK_URL_OCCALL
//        policy.CallbackBody = "key=$(key)&hash=$(etag)"
//        policy.PersistentOps = "avthumb/m3u8/segment/10"
//        policy.PersistentNotifyUrl = CALLBACK_URL_OCCALL
//
//    }
//
//    upToken := policy.Token(mac_occall)
//
//    var ret qiniuio.PutRet //用于存取返回的信息
//    var extra = &qiniuio.PutExtra{
//        //Params:    params,
//        //MimeType:  mieType,
//        //Crc32:     crc32,
//        //CheckCrc:  CheckCrc,
//    }
//
//    id := newId()
//    err := qiniuio.Put(nil, &ret, upToken, id.Hex(), data, extra)
//
//    if err != nil {
//        //上传产生错误
//        fmt.Errorf("upload img to qiniu server error: %v", err)
//        return nil, 500, fmt.Sprintf("upload img to qiniu server error: %v", err)
//    }
//
//    //上传成功，处理返回值
//    log.Print(ret.Hash, ret.Key)
//
//    resource := &Resource{
//        Id:     id,
//        Name:   filename,
//        Bucket: BUCKET_OCCALL,
//        Key:    ret.Key,
//        Hash:   ret.Hash,
//        Tp:     tp,
//        Status: RESOURCE_STATUS_VALID,
//        Ct:     tick(),
//    }
//
//    err = addResource(ds.se, resource)
//    if err != nil {
//        if dup(err) {
//            return nil, 400, "duplicate key error"
//        }
//        return nil, 500, fmt.Sprintf("create file resource error: %v", err)
//    }
//
//    //如果是视频文件，上传完成后进行转码操作，转成MP4文件
//    if resource.Tp == RESOURCE_TYPE_VIDEO {
//        //fmt.Println("encode video:", r.Key)
//        persistentId, err := encodeByOccall(resource.Key)
//        if err != nil {
//            fmt.Printf("encode video error: %v\n", err)
//            return nil, 500, fmt.Sprintf("encode video error: %v", err)
//        }
//
//        resource.PersistentId = persistentId
//        err = updateResource(ds.se, resource)
//        chk(err)
//    }
//
//    return resource, 0, ""
//}
//
//func qiniuImgUrlByOccall(id string) string {
//    if id == "" {
//        return ""
//    }
//    return fmt.Sprintf("http://%s/%s", QINIU_DOMAIN_OCCALL, id)
//}
//
//func originalImgUrlByOccall(id bson.ObjectId) string {
//    if id.Hex() == "" {
//        return ""
//    }
//    return fmt.Sprintf("http://%s/%s", QINIU_DOMAIN_OCCALL, id.Hex())
//}
//
//func fileDownloadByOccall(res *Resource) string {
//    return fmt.Sprintf("http://%s/%s?attname=%s", QINIU_DOMAIN_OCCALL, res.Id.Hex(), res.Name)
//}
//
//func videoUrlByOccall(id string) string {
//    if id == "" {
//        return ""
//    }
//    return fmt.Sprintf("http://%s/mp4_%s", QINIU_DOMAIN_OCCALL, id)
//}
//
//func videoPosterByOccall(id string) string {
//    if id == "" {
//        return ""
//    }
//    return fmt.Sprintf("http://%s/img_%s", QINIU_DOMAIN_OCCALL, id)
//}
//
//func qiniuOccallTransferHandler(ds *Ds) (int, string) {
//    resList, err := findResourceByQuery(ds.se, nil)
//    chk(err)
//    log.Printf("total: %v", len(resList))
//    for i, res := range resList {
//        log.Printf("dealing: %v", i+1)
//        imgUrl := qiniuImgUrl(res.Id.Hex())
//        //		log.Printf("imgUrl: %s", imgUrl)
//
//        if imgUrl == "" {
//            log.Printf("img url empty %v", i)
//            continue
//        }
//        resp, err := http.Get(imgUrl)
//        defer resp.Body.Close()
//        chk(err)
//        transferStart(ds, resp.Body, res, res.Tp)
//        url := qiniuImgUrlByOccall(res.Id.Hex())
//        log.Printf("t url: %s", url)
//
//    }
//
//    return 200, "ok"
//}
//
//func transferStart(ds *Ds, data io.Reader, res *Resource, tp int) (*Resource, int, string) {
//    // ret       变量用于存取返回的信息，详情见 io.PutRet
//    // uptoken   为业务服务器端生成的上传口令
//    // key       为文件存储的标识
//    // r         为io.Reader类型，用于从其读取数据
//    // extra     为上传文件的额外信息,可为空， 详情见 io.PutExtra, 可选
//
//    policy := &rs.PutPolicy{
//        Scope: BUCKET_OCCALL,
//    }
//
//    if tp == RESOURCE_TYPE_VIDEO {
//        policy.CallbackUrl = CALLBACK_URL_OCCALL
//        policy.CallbackBody = "key=$(key)&hash=$(etag)"
//        policy.PersistentOps = "avthumb/m3u8/segment/10"
//        policy.PersistentNotifyUrl = CALLBACK_URL_OCCALL
//
//    }
//
//    upToken := policy.Token(mac_occall)
//
//    var ret qiniuio.PutRet //用于存取返回的信息
//    var extra = &qiniuio.PutExtra{
//        //Params:    params,
//        //MimeType:  mieType,
//        //Crc32:     crc32,
//        //CheckCrc:  CheckCrc,
//    }
//
//    id := res.Id
//    err := qiniuio.Put(nil, &ret, upToken, id.Hex(), data, extra)
//
//    if err != nil {
//        //上传产生错误
//        fmt.Errorf("upload img to qiniu server error: %v", err)
//        return nil, 500, fmt.Sprintf("upload img to qiniu server error: %v", err)
//    }
//
//    //上传成功，处理返回值
//    log.Print(ret.Hash, ret.Key)
//
//    //    resource := &Resource{
//    //        Id:     id,
//    //        Name:   filename,
//    //        Bucket: BUCKET_OCCALL,
//    //        Key:    ret.Key,
//    //        Hash:   ret.Hash,
//    //        Tp:     tp,
//    //        Status: RESOURCE_STATUS_VALID,
//    //        Ct:     tick(),
//    //    }
//    //
//    //    err = addResource(ds.se, resource)
//    //    if err != nil {
//    //        if dup(err) {
//    //            return nil, 400, "duplicate key error"
//    //        }
//    //        return nil, 500, fmt.Sprintf("create file resource error: %v", err)
//    //    }
//
//    res.Bucket = BUCKET_OCCALL
//    res.Hash = ret.Hash
//    err = updateResource(ds.se, res)
//    chk(err)
//
//    //如果是视频文件，上传完成后进行转码操作，转成MP4文件
//    if res.Tp == RESOURCE_TYPE_VIDEO {
//        //fmt.Println("encode video:", r.Key)
//        persistentId, err := encodeByOccall(res.Key)
//        if err != nil {
//            fmt.Printf("encode video error: %v\n", err)
//            return nil, 500, fmt.Sprintf("encode video error: %v", err)
//        }
//
//        res.PersistentId = persistentId
//        err = updateResource(ds.se, res)
//        chk(err)
//    }
//
//    return res, 0, ""
//}
