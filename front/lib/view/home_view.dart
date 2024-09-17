import 'package:flutter/material.dart';

import '../service/vehicle.dart';
import 'vehicle_view.dart';

class HomeView extends StatefulWidget {
  const HomeView({super.key});

  @override
  State<HomeView> createState() => _HomeViewState();
}

class _HomeViewState extends State<HomeView> {
  final VehicleService _vehicleService = VehicleService();
  var idController = TextEditingController();
  var yearController = TextEditingController();
  var modelController = TextEditingController();
  var plateController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Adicionar novo veículo'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Padding(
              padding: const EdgeInsets.all(5.0),
              child: ElevatedButton(
                onPressed: () {
                  Navigator.push(
                    context,
                    MaterialPageRoute(
                        builder: (context) => const VehiclesView()),
                  );
                },
                child: const Text('Ir para listagem'),
              ),
            ),
            SizedBox(
              width: 200,
              child: TextFormField(
                controller: yearController,
                decoration: const InputDecoration(labelText: 'Ano'),
              ),
            ),
            SizedBox(
              width: 200,
              child: TextFormField(
                controller: modelController,
                decoration: const InputDecoration(labelText: 'Modelo'),
              ),
            ),
            SizedBox(
              width: 200,
              child: TextFormField(
                controller: plateController,
                decoration: const InputDecoration(labelText: 'Placa'),
              ),
            ),
            Padding(
              padding: const EdgeInsets.all(5.0),
              child: ElevatedButton(
                onPressed: () {
                  _vehicleService.postData(
                      year: int.tryParse(yearController.text) ?? 0,
                      model: modelController.text,
                      plate: plateController.text);
                  yearController.clear();
                  modelController.clear();
                  plateController.clear();
                },
                child: const Text('Salvar'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
