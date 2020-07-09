import math

import numpy
import pytesseract
from PIL import Image
from twisted.internet.process import detector

from google.test import binarizing


def rolingImage(img):
    open_cv_image = numpy.array(img)
    # Convert RGB to BGR
    open_cv_image = open_cv_image[:, :, ::-1].copy()
    result = detector.detect_faces(open_cv_image)
    x1 = result[0]['keypoints']['left_eye'][0]
    y1 = result[0]['keypoints']['left_eye'][1]
    x2 = result[0]['keypoints']['right_eye'][0]
    y2 = result[0]['keypoints']['right_eye'][1]

    xy = math.fabs(math.sqrt((x1 - x2) ** 2 + (y1 - y2) ** 2))
    h = math.fabs(y2 - y1)
    sino = h / xy
    o = math.degrees(math.asin(sino))
    o_fix = 0;
    if y2 - y1 < 0:
        o_fix = 0 - o
    else:
        o_fix = o

    print(xy, h, sino, o, o_fix)

    img_fix = img.rotate(o_fix, expand=True)
    return img_fix


def idcardFront(pic_path):
    dd = ['', '', '', '', '', '']
    try:
        img1 = Image.open(pic_path)
        # img1 = rolingImage(img1)
        # img2 = img1.convert('L')
        # img3 = binarizing(img2, 130)
        code = pytesseract.image_to_string(img1, lang='chi_sim')
        sa = code.split("\n")
        s = 1;
        for one in sa:
            print(s, one);
            s = s + 1
        # 身份证——姓名
        name = ""

        if (len(sa) >= 1):
            name_start = sa[0].find('名')
            if (name_start != -1):
                name = sa[0][name_start + 1:].replace(" ", "")
            else:
                name = sa[0].replace(" ", "")
        dd[0] = name
        print("name", name)

        sex = ""
        # 身份证——性别
        pma = sa[1]
        if (len(sa) >= 2):
            sex_start = sa[1].find('别')
            if (name_start != -1):
                sex = sa[1][sex_start + 1:].replace(" ", "")[0:1]
                pma = sa[1][sex_start + 1:].replace(" ", "")[1:]
        dd[1] = sex
        print("sex", sex)
        # 身份证——民族
        mingzu = ""
        if (len(sa) >= 2):
            mingzu_start = pma.find('族')
            if (mingzu_start != -1):
                mingzu = pma[mingzu_start + 1:].replace(" ", "")

        dd[2] = mingzu
        print("mingzu", mingzu)
        # 身份证——出生
        birthday = ""
        if (len(sa) >= 3):
            birthday_start = sa[2].find('生')
            if (birthday_start != -1):
                birthday = sa[2][birthday_start + 1:].replace(" ", "")
            else:
                birthday = sa[2].replace(" ", "")
            birthday_year = birthday.find('年')
            birthdayyear = ""
            birthdaymouth = ""
            birthdayday = ""
            if (birthday_year != -1):
                birthdayyear = birthday[0:birthday_year]
                birthday_mouth = birthday.find('月')

                if (birthday_mouth != -1):
                    birthdaymouth = birthday[birthday_year + 1:birthday_mouth]
                    if (len(birthdaymouth) == 1):
                        birthdaymouth = "0" + birthdaymouth
                    birthdayday = birthday[birthday_mouth + 1:len(birthday) - 1]
                    if (len(birthdayday) == 1):
                        birthdayday = "0" + birthdayday
                birthday = birthdayyear + birthdaymouth + birthdayday
            else:
                new_birth = filter(lambda ch: ch in '0123456789', birthday)
                birthday = ''.join(list(new_birth));

        dd[3] = birthday
        print("birthday", birthday)
        # 身份证——住址
        live = ""
        if (len(sa) >= 5):
            live_start = sa[4].find('址')
            if (live_start != -1):
                live = sa[4][live_start + 1:].replace(" ", "")
            else:
                live = sa[4].replace(" ", "")
        if (len(sa) >= 6):
            live += sa[5].replace(" ", "")

        dd[4] = live
        print("live", live)

        # 身份证——证件号
        w, h = img1.size
        region = (0, h / 8 * 6, w, h)
        # 裁切身份证号码图片
        cropImg = img1.crop(region)
        img2 = cropImg.convert('L')
        img3 = binarizing(img2, 105)
        # Image._show(img3)
        idcard = pytesseract.image_to_string(img3, lang='chi_sim')
        idcard_start = idcard.find('码')
        if (idcard_start != -1):
            idcard = idcard[idcard_start + 1:].replace(" ", "")
        else:
            idcard = idcard.replace(" ", "")
        new_crazy = filter(lambda ch: ch in '0123456789Xx', idcard)
        idcard = ''.join(list(new_crazy));
        dd[5] = idcard
        print("idcard", idcard)

    except Exception as e:
        print('Error:', e)
    return dd


def idcardBack(pic_path):
    dd = ['', '']
    try:
        img1 = Image.open(pic_path)
        w, h = img1.size
        region = (0, h / 5 * 3, w, h)
        # 裁切身份证号码图片
        cropImg = img1.crop(region)
        img2 = cropImg.convert('L')
        img3 = binarizing(img2, 100)
        # Image._show(img3)
        code = pytesseract.image_to_string(img3, lang='chi_sim')
        sa = code.split("\n")
        s = 1;
        for one in sa:
            print(s, one);
            s = s + 1
        # 身份证——签发机关
        guan = ""
        expering = ""
        for one in sa:
            print(s, one);
            s = s + 1
            guan_start = one.find('关')
            if (guan_start != -1):
                guan = one[guan_start + 1:].replace(" ", "")
                dd[0] = guan
            expering_start = one.find('限')
            if (expering_start != -1):
                expering = one[expering_start + 1:].replace(" ", "").replace("一", "-").replace("_", "-").replace("_",
                                                                                                                 "-").replace(
                    "O", "0").replace("o", "0")
                dd[1] = expering

        # 身份证——签发机关
        print("guan", guan)
        # 身份证——有效期

        print("expering", expering)
    except Exception as e:
        print('Error:', e)
    return dd

if __name__ == '__main__':

    # sk_back = idcardBack(filename_back)
    sk_front = idcardFront("1.jpeg")
    # sk_front = idcardFront("test1.png")

    ret = {}

    ret['name'] = sk_front[0]
    ret['sex'] = sk_front[1]
    ret['mingzu'] = sk_front[2]
    ret['birthday'] = sk_front[3]
    ret['live'] = sk_front[4]
    ret['idcard'] = sk_front[5]
    # ret['guan'] = sk_back[0]
    # ret['expering'] = sk_back[1]
    print(ret)







