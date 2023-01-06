import React,{useState,useEffect} from 'react'
import { Card, Col, Row, Tag, Avatar, Typography, Modal, Space} from 'antd';
import { UserOutlined,EditOutlined,DeleteOutlined } from '@ant-design/icons';
import { startCase} from 'lodash'
import InputForm from './InputForm';
import axios from 'axios';
const { Text }= Typography

function DisplayCard() {

  const [users,setUsers]= useState(null)
  const [userDetail,setUserDetail]=useState({})
  const [isModalOpen, setIsModalOpen] = useState(false);
  useEffect(()=>{
     fetchData();
  },[])
  const fetchData=async()=>{
    const responseData = await axios.get('http://127.0.0.1:4000/api/users/list')
    const {data =[]} = responseData.data
    console.log(data)
    setUsers(data)
  }
  const showModal = () => {
    setIsModalOpen(true);
  };
  const handleOk = () => {
    setUserDetail({})
    setIsModalOpen(false);
  };
  const handleCancel = () => {
    setUserDetail({})
    setIsModalOpen(false);
  };
  const getStatus=(status)=>{
      return status?<Tag color="#87d068">active</Tag>:<Tag color="#f50">inactive</Tag>
  }
  const getAvatar=(status)=>{
    return status?<Avatar style={{ backgroundColor: '#87d068' }} icon={<UserOutlined />} />:<Avatar style={{ backgroundColor: '"#f50"' }} icon={<UserOutlined />} />
  }
  const editCard=(_id,name,email,status,password)=>{
    setUserDetail((prevVal)=> ({...prevVal,_id,name,email,status,password}))
    console.log(userDetail)
    showModal();
  }
  const deleteCard=async(_id)=>{
    await axios.delete(`http://127.0.0.1:4000/api/users/delete/${_id}`)
  }
  return(
    <div className="site-card-wrapper">
    <Row gutter={16}>
      {
        users && users.map(({_id,name,email,status,password})=>{
        return(
          <Col span={8} style={{paddingBottom:'10px'}}>
            <Card 
              title={startCase(name)} 
              bordered={false} 
              extra={
              <Space size={10}>
                  <EditOutlined key="edit" style={{cursor:'pointer'}} onClick={()=>{editCard(_id,name,email,status,password)}}/>
                  <DeleteOutlined key="edit" style={{cursor:'pointer'}} onClick={()=>{deleteCard(_id)}}/>
              </Space>
              }>
                {getAvatar(status)}
                <p><Text strong>Name: </Text>{name}</p>
                <p><Text strong>Email: </Text>{email}</p>
                <p><Text strong>Status: </Text>{getStatus(status)}</p>
            </Card>
          </Col>)
        })}
    </Row>
    <Modal title="Update User" open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
        <InputForm 
        data={userDetail}
        type={"update"}
        />
      </Modal>
  </div>
  );
}
export default DisplayCard; 