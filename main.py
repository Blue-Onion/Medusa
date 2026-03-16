from ultralytics import YOLO

import cv2 
md=YOLO()
cap=cv2.VideoCapture(0)
while True:
    ret,frame=cap.read()
    print(frame)
    res=md(frame)
    print(res)
    annotated=res[0].plot()
    cv2.imshow("Hello",annotated)
    if cv2.waitKey(1)==27:
        break
